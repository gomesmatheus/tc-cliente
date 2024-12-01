Feature: ClienteUseCases

  Scenario: Registering a new valid client
    Given I have a valid client
    When I try to register the client
    Then the client should be registered successfully

  Scenario: Registering an invalid client
    Given I have an invalid client
    When I try to register the client
    Then the registration should fail with an error message "Cliente inválido"

  Scenario: Retrieving a client by valid CPF
    Given I have a valid CPF "12345678901"
    When I try to retrieve the client
    Then I should get the client details

  Scenario: Retrieving a client by invalid CPF
    Given I have an invalid CPF "987654321"
    When I try to retrieve the client
    Then I should receive an error "Cliente não encontrado"
