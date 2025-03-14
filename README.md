# What this tool does
This is a project that generates an easy to read description of the given codebase.

# Process

## Scanning
### Steps
1. Scans the filestructure
2. From the less to the most files that have dependencies.
3. Create a description yaml file for the file, containing minimal description of expected inputs and outputs and optional high level description of the logical process, also the yaml descriptions of the dependencies that are being used in the file, also keywords of the functionality.
4. If all the files of the folder have generated descriptions, creating yaml for the folder that describes what it contains, with keywords from the childen.
5. Recursevely this happens for the entire codebase

## RAG
### Steps
1. get keywords from the user prompt
2. recuresively get the relative descripitions docs of the files and folders that those keywords match
3. combine all in one text and copy them into the users clipboard directly to paste them into cursor