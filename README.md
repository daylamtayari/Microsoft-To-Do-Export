# Microsoft To Do Export

An unofficial solution to export all of your data such as lists, tasks, and attachments from Microsoft To Do.

## Features
- Export into the Todoist CSV import format
- Export all data into JSON
- Retrieve all lists and tasks, including completed tasks
- Includes notes and sub-tasks in exports
- Retrieve all or task-specific attachments
- Raw data output for piping into other tools
- Token can be specified either as an environment variable, in a file, or as an input parameter

### Developer Features
- `mstodo` package that handles all retrieval interactions for Microsoft To Do
- `todoistcsv` package that implements the Todoist CSV import format and all of its types
- `mstodo-to-todoistcsv` package that converts a Microsoft To Do list of tasks into a Todoist CSV import format

### Export Formats
- Todoist CSV
- JSON
- Generic CSV

**If there is an export format you want me to add, feel free to create an issue suggesting a format.**

## Installation

### Go Install

If you already have Go installed, you can install it directly by running:
```bash
go install github.com/daylamtayari/Microsoft-To-Do-Export@latest
```

### Binaries

Download the appropriate binary for your system (Windows, MacOS, or Linux) from the releases page: [https://github.com/daylamtayari/Microsoft-To-Do-Export/releases](https://github.com/daylamtayari/Microsoft-To-Do-Export/releases)

## Usage

```
$ Microsoft-To-Do-Export --help
Microsoft To Do Export

Usage:
  Microsoft-To-Do-Export [command]

Available Commands:
  attachment  Retrieve task attachment
  completion  Generate the autocompletion script for the specified shell
  export      Export task lists
  help        Help about any command
  list        List task lists
  version     Retrieve program version

Flags:
      --debug               Enable debug logging
  -h, --help                help for Microsoft-To-Do-Export
  -t, --token string        Token value
      --token-file string   File containing token

Use "Microsoft-To-Do-Export [command] --help" for more information about a command.
```

### Token Retrieval

A Microsoft Graph API token is required to use this program.

To retrieve it perform the following steps:
1. Go to Microsoft's Graph API Explorer: https://developer.microsoft.com/en-us/graph/graph-explorer
2. Sign in with the account that you want to retrieve the tasks from
3. Select the `my To Do task lists` option
4. Click the `Modify permissions` tab and consent to the `Tasks.ReadWrite` permission
5. Navigate to the `Access token` tab and copy the access token

You can pass your token any of the following ways:
- Pass it directly using the `--token YOURTOKEN` flag
- Store it in a file and specify the file in the `--token-file YOURTOKEN.FILE` flag
- Set it to be the `MSTODO_EXPORT_TOKEN` environment variable: `export MSTODO_EXPORT_TOKEN=YOURTOKEN`

### Export

```
# Example command:
$ Microsoft-To-Do-Export export --output myexport.csv --type todoist

$ Microsoft-To-Do-Export export --help
Export task lists

Usage:
  Microsoft-To-Do-Export export [flags]

Flags:
      --completed       Include completed tasks (NOTE: IF USING TODOIST EXPORT TYPE, TODOIST'S IMPORT FORMAT HAS NO WAY OF MARKING A TASK AS
 COMPLETED SO THIS WILL RESULT IN ALL TASKS MARKED AS UNCOMPLETED)
  -h, --help            help for export
  -o, --output string   Output file name (default "mstodo_export.{file_type}")
  -r, --raw             Output to stdout instead of a file and no table
      --type string     Output type (accepted values: 'json', 'todoist', 'csv') (default "json")

Global Flags:
      --debug               Enable debug logging
  -t, --token string        Token value
      --token-file string   File containing token

```

### Attachments
```bash
# Example command exporting of all attachments, including from completed tasks:
$ Microsoft-To-Do-Export attachment --all --completed --output todo_attachments

# Example command exporting all attachments from a single task:
$ Microsoft-To-Do-Export attachment --list LISTID --task TASKID --output task_attachments
```

## Help

For any assistance or feature requests, please create an issue on the GitHub page.

## License:

This project is licensed under GPL v3.0.  
The complete license: [LICENSE](https://github.com/daylamtayari/Microsoft-To-Do-Export/blob/master/LICENSE).  
For more details, please check out the official page: https://www.gnu.org/licenses/gpl-3.0.en.html  

## Disclaimer:

This project is in no way affiliated with Microsoft and Todoist, and their respective affiliates.
