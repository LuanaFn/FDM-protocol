Feature: Order
  Order products from end to end

  Scenario: Order a list of products from current vendor
    When I do not add products from external vendors
    And I do not provide a callback URL
    And I order a list of products from current vendor
    Then I receive confirmation that my order was sent
    And the order is sent to the order business service
