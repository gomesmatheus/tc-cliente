Feature: ClienteHandler

  Scenario: Creating a new client with valid data
    Given I have valid client data
    When I send a POST request to "/cliente"
    Then the response status should be 201
    And the response message should be "Cliente inserido"

  Scenario: Creating a new client with invalid data
    Given I have invalid client data
    When I send a POST request to "/cliente"
    Then the response status should be 400
    And the response message should be "400 bad request"

  Scenario: Retrieving a client with valid CPF
    Given I have a valid CPF "123"
    When I send a GET request to "/cliente/123"
    Then the response status should be 200
    And the response should contain the client information

  Scenario: Retrieving a client with invalid CPF
    Given I have an invalid CPF "invalid-cpf"
    When I send a GET request to "/cliente/invalid-cpf"
    Then the response status should be 400
    And the response message should be "Formato de CPF inválido"
