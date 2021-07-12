Feature: Hangup    
  In order to hangup a call after a certain time in seconds
  As an end user
  I want to have my call automatically off

Background: Setup
    Given my test setup runs

  Scenario: Hangup a call
    And "NumberA" configured to hangup after "3" seconds
    When I make a call from "NumberA" to "NumberB"
    Then "NumberA" should get last call duration greater than or equals to "3"