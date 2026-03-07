Device tracking

You can store things like:
device_id
ip_address
user_agent
last_seen

So users can see:
"Logged in devices"

Now the honest challenges 
Auth systems are extremely security sensitive.
You will eventually need:
security audits
Example problems:
token replay
timing attacks
password reset vulnerabilities
CSRF
JWT signing mistakes

Even big companies get auth wrong.

edge cases explode quickly
Examples:
device revocation
token rotation
session hijacking
email change verification
multi-factor auth
password reset flows
account lockouts

Auth systems grow very complex.