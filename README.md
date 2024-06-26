# Go dreaming of adventure sample

A developer sample written in Go that demonstrates Gemini's creative writing
abilities. With user input, Gemini writes a novella one section at a time.

<a href="https://idx.google.com/import?url=https://github.com/google-gemini/go-dreaming-of-adventure-sample">
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://cdn.idx.dev/btn/open_dark_32@2x.png">
  <source media="(prefers-color-scheme: light)" srcset="https://cdn.idx.dev/btn/open_light_32@2x.png">
  <img height="32" alt="Open in IDX" src="https://cdn.idx.dev/btn/open_purple_32@2x.png">
</picture>
</a>

## Environment setup

This sample app can be opened in _Project IDX_, or run in your local dev environment.

## Project IDX

1. Open this repo in Project IDX:
   - [Open in Project IDX](https://idx.google.com/import?url=https://github.com/google-gemini/go-dreaming-of-adventure-sample)
   - Wait for the import process to complete
   - Open the IDX Panel and click "Authenticate" with the Gemini API integration. 
   - Once authenticated, click to get a key which will be copied to your keyboard.
   - Add the key to the env variable section in `.idx/dev.nix`.

1. Open a new terminal window:
   - Open the command palette (CTRL/CMD-SHIFT-P)
   - Begin typing **terminal**
   - Select **Terminal: Create New Terminal**
   - Run `go run .`

## Local dev environment

1. Clone this repository: `git clone https://github.com/google-gemini/go-dreaming-of-adventure-sample`

1. Verify that Go 1.22 or later is installed:
   - Verify version with `go version`
   - In needed, install Go, see: https://go.dev/doc/install


## Run the sample

1. Get a Gemini API key
    - Launch Google AI Studio: https://aistudio.google.com/
    - Click **Get API Key**

1. Set the API Key in the `API_KEY` environment varaible
    - `export API_KEY=<your_api_key>`

1. Compile and run the program:
   - `go run .`

1. When asked "What do you want to dream about?", answer with something fun.
   - For example, type: `I want to dream about unicode`
