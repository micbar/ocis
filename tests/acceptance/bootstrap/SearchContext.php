<?php declare(strict_types=1);
/**
 * ownCloud
 *
 * @author Artur Neumann <artur@jankaritech.com>
 * @copyright Copyright (c) 2018 Artur Neumann artur@jankaritech.com
 *
 * This code is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License,
 * as published by the Free Software Foundation;
 * either version 3 of the License, or any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>
 *
 */

use Behat\Behat\Context\Context;
use Behat\Behat\Hook\Scope\BeforeScenarioScope;
use Behat\Gherkin\Node\TableNode;
use PHPUnit\Framework\Assert;
use TestHelpers\WebDavHelper;
use TestHelpers\HttpRequestHelper;

require_once 'bootstrap.php';

/**
 * context containing search related API steps
 */
class SearchContext implements Context {
	private FeatureContext $featureContext;

	/**
	 * @When user :user searches for :pattern using the WebDAV API
	 * @When user :user searches for :pattern and limits the results to :limit items using the WebDAV API
	 * @When user :user searches for :pattern using the WebDAV API requesting these properties:
	 * @When user :user searches for :pattern and limits the results to :limit items using the WebDAV API requesting these properties:
	 * @When user :user searches for :pattern inside folder :scope using the WebDAV API
	 * @When user :user searches for :pattern inside folder :scope in space :spaceName using the WebDAV API
	 *
	 * @param string $user
	 * @param string $pattern
	 * @param string|null $limit
	 * @param string|null $scope
	 * @param string|null $spaceName
	 * @param TableNode|null $properties
	 *
	 * @return void
	 */
	public function userSearchesUsingWebDavAPI(
		string    $user,
		string    $pattern,
		?string   $limit = null,
		?string   $scope = null,
		?string   $spaceName = null,
		TableNode $properties = null
	): void {
		// NOTE: because indexing of newly uploaded files or directories with ocis is decoupled and occurs asynchronously
		// short wait is necessary before searching
		sleep(5);
		$user = $this->featureContext->getActualUsername($user);
		$baseUrl = $this->featureContext->getBaseUrl();
		$password = $this->featureContext->getPasswordForUser($user);
		if (str_contains($pattern, '$')) {
			$date = explode("$", $pattern);
			switch ($date[1]) {
				case "today":
					$pattern = $date[0] . date('Y-m-d', strtotime('today'));
					break;
				case "yesterday":
					$pattern = $date[0] . date('Y-m-d', strtotime('yesterday'));
					break;
				default:
					throw new Exception("cannot convert the date");
			}
		}
		$body
			= "<?xml version='1.0' encoding='utf-8' ?>\n" .
			"	<oc:search-files xmlns:a='DAV:' xmlns:oc='http://owncloud.org/ns' >\n" .
			"		<oc:search>\n";
		if ($scope !== null) {
			$scope = \trim($scope, "/");
			$resourceID = $this->featureContext->spacesContext->getResourceId($user, $spaceName ?? "Personal", $scope);
			$pattern .= " scope:$resourceID";
		}
		$body .= "<oc:pattern>$pattern</oc:pattern>\n";
		if ($limit !== null) {
			$body .= "			<oc:limit>$limit</oc:limit>\n";
		}

		$body .= "		</oc:search>\n";
		if ($properties !== null) {
			$propertiesRows = $properties->getRows();
			$body .= "	<a:prop>";
			foreach ($propertiesRows as $property) {
				$body .= "<$property[0]/>";
			}
			$body .= "	</a:prop>";
		}
		$body .= "	</oc:search-files>";
		$davPathVersionToUse = $this->featureContext->getDavPathVersion();
		$davPath = WebDavHelper::getDavPath($doDavRequestAsUser ?? $user, $davPathVersionToUse, 'files');

		if ($davPathVersionToUse == WebDavHelper::DAV_VERSION_NEW) {
			// Removes the last component('username' in this case) from the WebDAV path by going up one level in the directory structure.
			// e.g. remote.php/dav/files/Alice ==> remote.php/dav/files/
			$davPath = \dirname($davPath, 1);
		}

		$fullUrl = WebDavHelper::sanitizeUrl("$baseUrl/$davPath");
		$response = HttpRequestHelper::sendRequest(
			$fullUrl,
			$this->featureContext->getStepLineRef(),
			'REPORT',
			$user,
			$password,
			null,
			$body
		);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @Then file/folder :path in the search result of user :user should contain these properties:
	 *
	 * @param string $path
	 * @param string $user
	 * @param TableNode $properties
	 *
	 * @return void
	 * @throws Exception
	 */
	public function fileOrFolderInTheSearchResultShouldContainProperties(
		string    $path,
		string    $user,
		TableNode $properties
	): void {
		$user = $this->featureContext->getActualUsername($user);
		$this->featureContext->verifyTableNodeColumns($properties, ['name', 'value']);
		$properties = $properties->getHash();
		$fileResult = $this->featureContext->findEntryFromSearchResponse(
			$path
		);
		Assert::assertNotFalse(
			$fileResult,
			"could not find file/folder '$path'"
		);
		foreach ($properties as $property) {
			$property['value'] = $this->featureContext->substituteInLineCodes(
				$property['value'],
				$user
			);
			$fileResultProperty = $fileResult->xpath("d:propstat//" . $property['name']);
			if ($fileResultProperty) {
				Assert::assertMatchesRegularExpression(
					"/" . $property['value'] . "/",
					\trim((string)$fileResultProperty[0])
				);
				continue;
			}
			throw new Error("Could not find property '" . $property['name'] . "'");
		}
	}

	/**
	 * This will run before EVERY scenario.
	 * It will set the properties for this object.
	 *
	 * @BeforeScenario
	 *
	 * @param BeforeScenarioScope $scope
	 *
	 * @return void
	 */
	public function before(BeforeScenarioScope $scope): void {
		// Get the environment
		$environment = $scope->getEnvironment();
		// Get all the contexts you need in this context
		$this->featureContext = $environment->getContext('FeatureContext');
	}

	/**
	 * @Then /^the search result should contain these (?:files|entries) with highlight on keyword "([^"]*)"/
	 *
	 * @param TableNode $expectedFiles
	 * @param string $expectedContent
	 *
	 * @return void
	 *
	 * @throws Exception
	 */
	public function theSearchResultShouldContainEntriesWithHighlight(
		TableNode $expectedFiles,
		string    $expectedContent
	): void {
		$this->featureContext->verifyTableNodeColumnsCount($expectedFiles, 1);
		$elementRows = $expectedFiles->getRows();
		$foundEntries = $this->featureContext->findEntryFromSearchResponse(
			null,
			true
		);
		foreach ($elementRows as $expectedFile) {
			$filename = $expectedFile[0];
			$content = $foundEntries[$filename];
			// Extract the content between the <mark> tags
			preg_match('/<mark>(.*?)<\/mark>/s', $content, $matches);
			$actualContent = $matches[1] ?? '';

			// Remove any leading/trailing whitespace for comparison
			$actualContent = trim($actualContent);
			Assert::assertEquals(
				$expectedContent,
				$actualContent,
				"Expected text highlight to be '$expectedContent' but found '$actualContent'"
			);
		}
	}
}
