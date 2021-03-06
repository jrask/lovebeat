Alerters
========

You can setup lovebeat to send mails or issue outgoing webhooks (HTTP POST) to
your web service whenever a view changes state. This is done on a view by
modifying the configuration file.

Send Mail
---------

The first step is to specify the SMTP server address and the e-mail address
that will be used when sending the e-mails. It doesn't currently support
SMTP authentication, so you might want to run a local SMTP server to proxy
the sent e-mails.

The configuration file should look as following:

.. code-block:: ini

    [mail]
    server = "localhost:25"
    from = "lovebeat@example.com"

**TODO:** Write detailed configuration

Outgoing Webhooks
-----------------

When a view changes state, a POST will be sent to the URL(s) specified in the
configuration. The JSON data that is sent follows:

.. code-block:: http

    POST /your/url/endpoint HTTP/1.1
    Content-Type: application/json
    Accept: application/json
    User-Agent: Lovebeat
    X-Lovebeat: 1

    {
      "name": "view.name.here",
      "from_state": "ok",
      "to_state": "error",
      "incident_number": 4
    }

The incident number is a monotonically incrementing counter that increases every
time a view transitions between **OK** and either **WARNING** or **ERROR**.

Slack
-----

Lovebeat can post messages to a slack_ channel whenever a view changes state.
First of all, setup an incoming webhook to get a Webhook URL that you will
enter in the lovebeat configuration file.

A working example would look like:

.. code-block:: ini

    [views.example]
    regexp = ".*"
    alerts = ["message-to-ops"]

    [alerts.message-to-ops]
    slack_channel = "#ops"

    [slack]
    webhook_url = "https://hooks.slack.com/services/T12345678/B12345678/abrakadabra"

Script
------

Lovebeat can run arbitrary scripts (or other executable files) whenever a view
changes state. The details of the alert will be posted as environment variables:

  * LOVEBEAT_VIEW=<name of the view>
  * LOVEBEAT_STATE=<the current state>
  * LOVEBEAT_PREVIOUS_STATE=<the previous state>
  * LOVEBEAT_INCIDENT=<incident number>

The script will also inherit any environment variables that Lovebeat was started
with.

The script's stdout and stderr will be printed, and the script will be invoked
with no arguments. If a script doesn't finish within 10 seconds, it will be
terminated. Remember to make your script executable using
``chmod a+x script.sh``.

Example of the configuration file:

.. code-block:: ini

    [views.example]
    regexp = ".*"
    alerts = ["test-alert"]

    [alerts.test-alert]
    script = "/path/to/script.sh"

The script (/path/to/script.sh) could look like:

.. code-block:: bash

    #!/bin/bash

    echo "Hello World"
    env

The output would then be (among other environment variables):

.. code-block:: text

    2016/01/26 18:10:56 INFO VIEW 'example', 11: state ok -> warning
    2016/01/26 18:10:56 INFO Running alert script /path/to/script.sh
    Hello World
    LOVEBEAT_VIEW=slack
    LOVEBEAT_STATE=WARNING
    LOVEBEAT_PREVIOUS_STATE=OK
    LOVEBEAT_INCIDENT=11

.. _slack: https://slack.com/
