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
  - generate ![image](https://github.com/akifkadioglu/organizer/assets/68117188/094f765f-67ef-4af3-9caa-07fa228f6dab)
  - list ![image](https://github.com/akifkadioglu/organizer/assets/68117188/945ac11d-f4b4-45d8-b74b-531b96c0dbf3)
  - get ![image](https://github.com/akifkadioglu/organizer/assets/68117188/461e4cb8-2407-4395-9f75-7ec915381922)
  - remove ![image](https://github.com/akifkadioglu/organizer/assets/68117188/85c83c5e-89e8-47ad-aea0-e623313e8756)
  - download as JSON ![image](https://github.com/akifkadioglu/organizer/assets/68117188/69c251db-8e8b-4ab5-b7c7-85713a78176f)





### Contributing

If you would like to contribute to this project, please fork the repository on GitHub and submit a pull request.

### Document

You can use Godoc for documentation, please run `make doc` and then go [here](http://localhost:6060).

### Thank you

- https://github.com/fatih
- https://github.com/spf13
- https://github.com/olekukonko
