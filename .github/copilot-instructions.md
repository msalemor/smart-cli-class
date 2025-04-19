mytool

The CLI application generates Kubernetes kubectl commands using AI based on the user's prompt.

Technologies used:
Go
Cobra
Microsoft OpenAI with an API KEY

The AI generates the following JSON:
```json
{
  "commands":[
    {"command":"kubectl","args":["get","all","-A"],"explanation":"List all nodes in all namespaces"}
  ]
}
```

# Golang CLI Best Practices with Cobra

## Project Structure

- Organize your CLI with a clear directory structure
    ```
    main.go
    /cmd           # Application entrypoints
    /internal      # Private application code
    /pkg           # Public library code
    /api           # API specifications
    ```
- Use `cmd/root.go` for the root command and subcommand directories

## Cobra Implementation

- Initialize with `cobra init` or `cobra-cli init`
- Create commands with descriptive names and documentation
    ```go
    var rootCmd = &cobra.Command{
        Use:   "mytool",
        Short: "A brief description of your application",
        Long:  `A longer description of your application`,
    }
    ```
- Use `PersistentFlags` for global flags and `Flags` for command-specific flags

## Command Design

- Follow UNIX philosophy: do one thing well
- Use subcommands for related functionality
- Keep command names consistent (verb-noun format)
- Implement "help" thoroughly for all commands

## User Experience

- Provide sensible default values
- Support both short and long flag names
- Use environment variables as fallbacks with `viper`
- Implement progress feedback for long-running operations

## Error Handling

- Return appropriate exit codes
- Display user-friendly error messages
- Log detailed errors for debugging
- Support different verbosity levels

## Configuration

- Use `viper` for configuration management
- Support multiple config formats (YAML, JSON, TOML)
- Look for config in standard locations
- Allow specifying config file location

## Testing

- Write unit tests for command logic
- Use table-driven tests for flag parsing
- Create integration tests for end-to-end workflows
- Mock external dependencies

## Distribution

- Use Go modules for dependency management
- Provide installation instructions
- Create releases for multiple platforms
- Use goreleaser for simplified distribution

## Documentation

- Include a README with installation and usage instructions
- Add examples for common use cases
- Generate man pages with `cobra`
- Implement shell completions