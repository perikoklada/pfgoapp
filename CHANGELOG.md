# 2024/03/29

Initial version.

# 2024/04/24

Bug fixes:
* Fix to correctly report the number of messages generated over a session
* Fix for the increased memory usage observed after a number of sessions
* Fix for the security issue that could lead to a CSRF attack

Features:
* HEX string generation and relevant tests were implemented
* Addition of the generated string to the message to the client was enabled
* Addition of a client app that generates multiple simultaneous sessions