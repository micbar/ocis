@skipOnReva
Feature: get shares filtered by type (user, group etc)
  As a user
  I want to filter the shares that I have received of a particular type (user, group etc)
  So that I can know about the status of the shares I've received

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |
    And group "grp1" has been created
    And user "Brian" has been added to group "grp1"
    And user "Alice" has created folder "/folderToShareWithUser"
    And user "Alice" has created folder "/folderToShareWithGroup"
    And user "Alice" has created folder "/folderToShareWithPublic"
    And user "Alice" has uploaded file with content "file to share with user" to "/fileToShareWithUser.txt"
    And user "Alice" has uploaded file with content "file to share with group" to "/fileToShareWithGroup.txt"
    And user "Alice" has uploaded file with content "file to share with public" to "/fileToShareWithPublic.txt"
    And user "Alice" has shared folder "/folderToShareWithUser" with user "Brian"
    And user "Alice" has shared folder "/folderToShareWithGroup" with group "grp1"
    And user "Alice" has created a public link share with settings
      | path        | /folderToShareWithPublic |
      | permissions | read                     |
      | password    | %public%                 |
    And user "Alice" has shared file "/fileToShareWithUser.txt" with user "Brian"
    And user "Alice" has shared file "/fileToShareWithGroup.txt" with group "grp1"
    And user "Alice" has created a public link share with settings
      | path        | /fileToShareWithPublic.txt |
      | permissions | read                       |
      | password    | %public%                   |


  Scenario Outline: getting shares shared to users
    Given using OCS API version "<ocs_api_version>"
    When user "Alice" gets the user shares shared by her using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And exactly 2 files or folders should be included in the response
    And folder "/Shares/folderToShareWithUser" should be included in the response
    And file "/Shares/fileToShareWithUser.txt" should be included in the response
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: getting shares shared to groups
    Given using OCS API version "<ocs_api_version>"
    When user "Alice" gets the group shares shared by her using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And exactly 2 files or folders should be included in the response
    And folder "/Shares/folderToShareWithGroup" should be included in the response
    And folder "/Shares/fileToShareWithGroup.txt" should be included in the response
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: getting shares shared to public links
    Given using OCS API version "<ocs_api_version>"
    When user "Alice" gets the public link shares shared by her using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And exactly 2 files or folders should be included in the response
    And folder "/folderToShareWithPublic" should be included in the response
    And folder "/fileToShareWithPublic.txt" should be included in the response
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: getting shares shared to users and groups
    Given using OCS API version "<ocs_api_version>"
    When user "Alice" gets the user and group shares shared by her using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And exactly 4 files or folders should be included in the response
    And folder "/Shares/folderToShareWithUser" should be included in the response
    And file "/Shares/fileToShareWithUser.txt" should be included in the response
    And folder "/Shares/folderToShareWithGroup" should be included in the response
    And folder "/Shares/fileToShareWithGroup.txt" should be included in the response
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |
