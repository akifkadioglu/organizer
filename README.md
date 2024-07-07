# Organizer

Organizer is a command-line tool that classifies files based on their mimetype. This tool categorizes and organizes files in a specified directory according to their types.

### Installation

To use the Organizer, you need to have Go installed on your computer.

1. Clone this repository from [GitHub](https://github.com/akifkadioglu/docs_organizer): `git clone https://github.com/akifkadioglu/organizer`
2. Navigate to the cloned directory: `cd organizer`
3. Build and install the Organizer: `make install`
4. Start using the CLI: `./organizer <command> <parameters>`

### Usage

- organizer categorize --path=<directory_path>: Categorizes files according to their types.
- organizer password --app=<app_name> --username=<password_name> generate: Generate a password and save it for you
  - generate ![image](https://github.com/akifkadioglu/organizer/assets/68117188/dc642e31-ffa9-4d18-bcf0-fe374fcfff50)

  - list ![image](https://github.com/akifkadioglu/organizer/assets/68117188/a1c479fe-1f39-4979-b208-303b4a42087d)

  - get ![image](https://github.com/akifkadioglu/organizer/assets/68117188/f57ea950-8b03-42cb-bcb9-e60feb7d0484)

  - remove ![image](https://github.com/akifkadioglu/organizer/assets/68117188/9fcaf5d0-d9b0-4a00-9ebd-7d06467406d7)

  - download as JSON ![image](https://github.com/akifkadioglu/organizer/assets/68117188/a1b98d21-1e55-47c5-a3de-2fa73832ece3)






### Contributing

If you would like to contribute to this project, please fork the repository on GitHub and submit a pull request.

### Document

You can use Godoc for documentation, please run `make doc` and then go [here](http://localhost:6060).

### Thank you

- https://github.com/fatih
- https://github.com/spf13
- https://github.com/olekukonko
