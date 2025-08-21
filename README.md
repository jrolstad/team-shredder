# Team Shredder

A Go-based CLI service for automated data management and cleanup operations across Atlassian products (Jira and Confluence). This tool provides a configurable, extensible framework for executing data deletion and archival actions based on customizable queries and criteria.

## 🚀 Features

- **Multi-Product Support**: Works with both Jira and Confluence instances
- **Flexible Query System**: Uses native Atlassian query languages (JQL for Jira, CQL for Confluence)
- **Configurable Actions**: Support for delete and archive operations
- **Organization-Based Processing**: Handles multiple organizations with separate configurations
- **Comprehensive Reporting**: Detailed execution results with success/failure tracking
- **Environment-Based Secrets**: Secure credential management through environment variables

## 📋 Prerequisites

- Go 1.23.2 or higher
- Atlassian API credentials (username and API key)
- Access to target Jira and/or Confluence instances

## 🛠️ Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/jrolstad/team-shredder.git
cd team-shredder

# Build the CLI
go build -o team-shredder ./cmd/cli

# Make executable (Unix-like systems)
chmod +x team-shredder
```

### Using Go Install

```bash
go install github.com/jrolstad/team-shredder/cmd/cli@latest
```

## ⚙️ Configuration

### Environment Variables

Set the following environment variables for authentication:

```bash
export ATLASSIAN_USERNAME="your-email@domain.com"
export ATLASSIAN_API_KEY="your-api-key"
```

### Configuration Repository

The application uses a configuration repository pattern to define data actions. Currently, it uses an in-memory repository with hardcoded examples. You can extend this to support:

- Database storage
- Configuration files (JSON, YAML)
- External configuration services

Example configuration structure:

```go
{
    "id": "1",
    "organizationId": "cf35573a-88ed-4070-a8fa-edbb5d42bb55",
    "appType": "confluence",
    "action": "delete",
    "site": "https://your-instance.atlassian.net/wiki",
    "query": "lastModified < now(\"-30d\") AND type = page"
}
```

## 🚀 Usage

### Basic Execution

```bash
# Run with default configuration
./team-shredder

# Or if installed via go install
team-shredder
```

### Example Output

```
----------------------
Org Id: cf35573a-88ed-4070-a8fa-edbb5d42bb55
  Site: https://jrolstad-sandbox-1.atlassian.net/wiki
  App Type: confluence
  Action: purgeTrash
    2025-01-15 10:30:00 => 2025-01-15 10:32:15
    Affected Items: 25
    Failures: 0
```

## 🏗️ Architecture

### Core Components

- **Orchestrators**: Coordinate the execution of data actions across organizations
- **Processors**: Handle specific product integrations (Jira, Confluence)
- **Repositories**: Manage configuration and data persistence
- **Services**: Provide shared functionality (secrets, environment)

### Supported Actions

#### Jira
- `delete`: Permanently remove issues
- `archive`: Archive issues (if supported by instance)

#### Confluence
- `delete`: Remove content pages
- `archive`: Archive content (if supported by instance)

### Query Languages

- **Jira**: Uses JQL (Jira Query Language)
- **Confluence**: Uses CQL (Confluence Query Language)

## 🧪 Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Building

```bash
# Build for current platform
go build -o team-shredder ./cmd/cli

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o team-shredder-linux ./cmd/cli
GOOS=windows GOARCH=amd64 go build -o team-shredder.exe ./cmd/cli
```

### Project Structure

```
team-shredder/
├── cmd/
│   └── cli/                 # CLI application entry point
├── internal/
│   └── pkg/
│       ├── core/            # Core utilities and helpers
│       ├── models/          # Data models and structures
│       ├── orchestrators/   # Business logic orchestration
│       ├── processors/      # Product-specific processors
│       ├── repositories/    # Data access layer
│       └── services/        # Shared services
├── .github/
│   └── workflows/           # CI/CD pipelines
└── README.md
```

## 🔧 Extending the Application

### Adding New Processors

1. Implement the `DataActionProcessor` interface
2. Register the processor in `DataActionProcessorFactory`
3. Add configuration support for the new product type

### Adding New Actions

1. Extend the processor's `Process` method
2. Implement the specific action logic
3. Update configuration models if needed

### Adding New Configuration Sources

1. Implement the `DataActionConfigurationRepository` interface
2. Update the factory or dependency injection to use the new repository

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

**Use with caution!** This tool performs destructive operations on your Atlassian data. Always:

- Test configurations in a sandbox environment first
- Review queries carefully before execution
- Ensure you have proper backups
- Verify you have the necessary permissions

## 🆘 Support

For issues, questions, or contributions:

- Open an issue on GitHub
- Review existing issues and discussions
- Check the code examples in the repository

---

**Note**: This tool is designed for administrative and maintenance tasks. Always ensure compliance with your organization's data retention policies and regulatory requirements.
