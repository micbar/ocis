Feature: propagation of etags when deleting a file or folder
  As a client app
  I want metadata (etags) of parent folders to change when a file or folder is deleted
  So that the client app can know to re-scan and sync the content of the folder(s)

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files
    And user "Alice" has created folder "/upload"


  Scenario Outline: deleting a file changes the etags of all parents
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/sub/file.txt"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    When user "Alice" deletes file "/upload/sub/file.txt" using the WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /upload     |
      | Alice | /upload/sub |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |

  @issue-4251
  Scenario Outline: deleting a folder changes the etags of all parents
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has created folder "/upload/sub/toDelete"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    When user "Alice" deletes folder "/upload/sub/toDelete" using the WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /upload     |
      | Alice | /upload/sub |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |


  Scenario Outline: deleting a folder with content changes the etags of all parents
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has created folder "/upload/sub/toDelete"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/sub/toDelete/file.txt"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    When user "Alice" deletes folder "/upload/sub/toDelete" using the WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path        |
      | Alice | /           |
      | Alice | /upload     |
      | Alice | /upload/sub |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |

  @skipOnReva
  Scenario Outline: sharee deleting a file changes the etags of all parents for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/sub/file.txt"
    And user "Alice" has shared folder "/upload" with user "Brian"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/upload"
    And user "Brian" has stored etag of element "/Shares/upload/sub"
    When user "Brian" deletes file "/Shares/upload/sub/file.txt" using the WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path               |
      | Alice | /                  |
      | Alice | /upload            |
      | Alice | /upload/sub        |
      | Brian | /                  |
      | Brian | /Shares            |
      | Brian | /Shares/upload     |
      | Brian | /Shares/upload/sub |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

  @skipOnReva
  Scenario Outline: sharer deleting a file changes the etags of all parents for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/sub/file.txt"
    And user "Alice" has shared folder "/upload" with user "Brian"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/upload"
    And user "Brian" has stored etag of element "/Shares/upload/sub"
    When user "Alice" deletes file "/upload/sub/file.txt" using the WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path               |
      | Alice | /                  |
      | Alice | /upload            |
      | Alice | /upload/sub        |
      | Brian | /                  |
      | Brian | /Shares            |
      | Brian | /Shares/upload     |
      | Brian | /Shares/upload/sub |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

  @issue-4251 @skipOnReva
  Scenario Outline: sharee deleting a folder changes the etags of all parents for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has created folder "/upload/sub/toDelete"
    And user "Alice" has shared folder "/upload" with user "Brian"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/upload"
    And user "Brian" has stored etag of element "/Shares/upload/sub"
    When user "Brian" deletes folder "/Shares/upload/sub/toDelete" using the WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path               |
      | Alice | /                  |
      | Alice | /upload            |
      | Alice | /upload/sub        |
      | Brian | /                  |
      | Brian | /Shares            |
      | Brian | /Shares/upload     |
      | Brian | /Shares/upload/sub |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

  @issue-4251 @skipOnReva
  Scenario Outline: sharer deleting a folder changes the etags of all parents for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has created folder "/upload/sub/toDelete"
    And user "Alice" has shared folder "/upload" with user "Brian"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/sub"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/upload"
    And user "Brian" has stored etag of element "/Shares/upload/sub"
    When user "Alice" deletes folder "/upload/sub/toDelete" using the WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path               |
      | Alice | /                  |
      | Alice | /upload            |
      | Alice | /upload/sub        |
      | Brian | /                  |
      | Brian | /Shares            |
      | Brian | /Shares/upload     |
      | Brian | /Shares/upload/sub |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

  @issue-4251
  Scenario Outline: deleting a file in a publicly shared folder changes its etag for the sharer
    Given using <dav-path-version> DAV path
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has created a public link share with settings
      | path        | upload   |
      | permissions | change   |
      | password    | %public% |
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    When the public deletes file "file.txt" from the last public link share using the password "%public%" and new public WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path    |
      | Alice | /       |
      | Alice | /upload |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |

  @issue-4251
  Scenario Outline: deleting a folder in a publicly shared folder changes its etag for the sharer
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has created a public link share with settings
      | path        | upload   |
      | permissions | change   |
      | password    | %public% |
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    When the public deletes folder "sub" from the last public link share using the password "%public%" and new public WebDAV API
    Then the HTTP status code should be "204"
    And these etags should have changed:
      | user  | path    |
      | Alice | /       |
      | Alice | /upload |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |
