# Seppuku

Seppuku is a tool for requesting self-termination of an EC2 instance in an autoscale group.

Running the seppuku command on an EC2 instance will set the instance health status to Unhealthy,
causing it to be destroyed and replaced according to autoscale policy.
