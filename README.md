# Email Sender Tool

This command-line tool sends emails using a provided SMTP server.

## Usage

Build the Docker image:

```bash
docker build -t email-sender .
```

Run the Docker container:

```bash
docker run -it --rm \
    -v /path/to/attachement:/app/path/to/attachement \
    email-sender \
    --from="Sender Name <sender@example.com>" \
    --to="Recipient 1 <recipient1@example.com>,recipient2@example.com" \
    --cc="CC 1 <cc1@example.com>,cc2@example.com" \
    --subject="Test Email" \
    --body="This is a test email" \
    --server="smtp.example.com" \
    --port=587 \
    --user="username" \
    --password="password" \
    --attachment="/app/path/to/attachment1.txt,/app/path/to/attachment2.txt"
```

