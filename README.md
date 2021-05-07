# Vulcan CLI

This is a simple CLI that creates a new project using a template project.

## Install

### Mac:

If you are using Mac OS you can download the binary with homebrew by executing the command `brew install caioeverest/tap/vulcan`

### Others

To install the app you can download the last version on the release page and put the binary at your path or bin folder.

> On future iterations we will add support for snap and flatpak

## Config

The first time you run the app it will ask you to add the basic configurations such as your name, email, and default ssh-key. But you can set other configurations like:

But you can manage your configurations by using the command:

`vulcan config`

flag         |required |description
-------------|---------|-------------
-l or --list |false    |list your current configs
--name       |false    |add or change the name setted
--email      |false    |add or change the email setted
--ssh        |false    |add or change the ssh setted

### Templates:

Templates are the reference files and folder that vulcan will use to create your new project.

To add and manage templates you can use the command

`vulcan config template`

flag            |required |description
----------------|---------|-------------
-l or --list    |false    |list your current stored templates
-d or --delete  |false    |delete a given template by name
-n or --name    |false    |will insert a new template with this name
-a or --address |false    |will insert a new template with this address

> If you don't pass any flag the command will start on insertion mode

> You can set a local template -- setting a reference project folder on your machine -- by making the address the project path like `path/to/template/` or use a remote template by using the git ssh address like `git@gitsomething.com:someowner/your-template.git` and if you want to use a specific branch `git@gitsomething.com:someowner/your-template.git;your-branch-name`

### Global custom placeholders:

Placeholders are the variables that you can set on your template and will be replaced by their value.

To add and manage global custom placeholders you can execute the following command:

`vulcan config placeholder`

flag           |required |description
---------------|---------|-------------
-l or --list   |false    |list your current stored placeholders
-d or --delete |false    |delete a given placeholder by name
-n or --name   |false    |will insert a new placeholder with this name
-v or --value  |false    |will insert a new placeholder with this value

> If you don't pass any flag the command will start on insertion mode

## Usage

### Global flags

flag          |required |description
--------------|---------|-------------
-d or --debug |false    |will start the debug mode

### Create

To create a new project you can run the following command

`vulcan create`

flag                 |required |description
---------------------|---------|-------------
-t or --template     |false    |will use this template to build your project
-p or --placeholders |false    |expect a JSON file with the custom placeholders that you can use on your project

## Templates

You can see an example on ``

## Placeholder

Placeholders are the variable that you will set on your templates to be replaced by values. The basic placeholder that all projects on vulcan have are:

- RepositoryName
- ProjectName
- Description
- Author
- Email

And as you can see in the example they can be recovered by using the form `{{ .PlaceholderName }}`

For custom placeholder, you can set on the global scope by adding the pairs on your config or using the flag -p or --placeholder pointing to a JSON file that will be read and loaded in your project. In this case, you **must** use the Custom prefix on your template. e.g. `{{ .Custom.something }}`.

