# smart-cli-class

Building a smart command-line interface application using OpenAI and Go

## My rules

- Always start in the playground
- Calling an LLM is easy what is difficult is the everything else
- He or she who can stuff the prompt gets the riches
- In prompt engineering a prompt can be though of as a semantic program

## Foundational concepts

- Calling an LLMs with REST
- Context window
- Instruct vs chat completion model
- Chat conversation management
  - Message types:
    - system
    - user
    - assistant
- OpenAI text vs json_mode

## Requirements

- Go knowledge
- OpenAI endpoint and key

## What are we building?

mytool az -p "List the running nodes"

## Steps we will follow today

###  Playground
- We will start in the playground
  - Open M365 copilot and type:
```text
system:
You are an AI that can help generate Kubernetes commands using kubectl based on the user's question or statement. Generate one or more commands depending on the ask.
Example:
{
  "commands": [ {"command":"kubectl", "args": ["get","nodes","-A"], "explanation":""//explanation of the command}],
}
Repond in JSON format. No epilogue or prologue.
user:
List all pods in all the namespaces.
```

### Project structure

Create the following structure:
```bash
mkdir mytool
mkdir -p mytool/cmd
mkdir -p mytool/pkg
cd mytool
touch main.go
touch cmd/root.go
touch pkg/openai.go
```

### Create the go module

Create a go module
```bash
cd mytool && go mod init mytool
```

### Create the cobra CLI

- with cai: `cai scaffold -p "Create a Go cobra cli called mytool. Create a root command (cmd/rootcmd.go) in the cmd folder, and run the command from main.go. Create subcommand called cmd/azcmd.go to run kubernetes commands and add it to root command. The cli should require one parameter called prompt for all subcommands. Create a pkg folder and create three files. One should create a pkg/types.go file to have a request (prompt:string) and response (commands:[]commands,command{command:string, args:[]string, explanation:string}) objects. The other file should be pkg/openai.go to call a chat completion using a rest post command with the request and response objects. Call the third file, pkg/settings.go. This file should be able to read openai.json file using godotenv, create a structure (endpoint:string,apikey:string,system_prompt:string,prompt:string), and expose the settings a a singleton."`

- Create a Cobra CLI shell
- A structure to receive the command to be executed
- A function to call a chat completion
- A function to exexute the commands
