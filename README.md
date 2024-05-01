# Tracky

Tracky is a simple tool that identifies pairs of developers who frequently contribute to the same files in a git
repository.

## Prerequisites

- Go 1.22 or later

## Installation

1. Clone the repository to your local machine:

```bash
git clone https://github.com/robinsmith-source/tracky.git
```

2. Navigate to the project directory:

```bash
cd tracky
```

3. Build the application:

```bash
go build -o tracky
```

This will create an executable file named `tracky` (or `tracky.exe` on Windows).

## Usage

First ensure that the `tracky` executable is in your system's PATH.

To do this on Unix-based systems, run the following command:

```bash
export PATH=$PATH:/path/to/tracky
```

To do this on Windows, add the path to the `tracky` executable to the system's PATH environment variable.

Then navigate to the root of the git repository you want to analyze and run the following command:

```bash
./tracky
```

This will output a list of pairs of developers who frequently contribute to the same files in the repository.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
