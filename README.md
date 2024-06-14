# ASCII Art Output

## Project Description

ASCII Art Web is a web based tool developed in Go and HTML that converts text input into ASCII art using predefined character sets and displays it on the browser. This tool also allows users choose from multiple ASCII art styles to enhance the visual representation of their text.

### Objectives

The primary objective of this project is to provide a flexible and user-friendly tool for generating ASCII art from text input. It offers the following features:

- Conversion of text input into ASCII art using customizable character sets.
- Multiple ASCII art styles to choose from, including standard, shadow, and thinkertoy.


<!-- TABLE OF CONTENTS -->
<details>
  <summary style="font-weight: bold; font-size: 1.4em;" >Table of Contents</summary>
  <ol>
    <li>
      <a href="#implementation-details">Implementation details</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#optional-ascii-art-styles">Optional ASCII Art Styles</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#authors">Authors</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


## Implementation details

ASCII Art Web is designed to provide users with a simple yet powerful tool for creating ASCII art from text input. It utilizes a set of ASCII characters to represent text in a visually appealing manner. The project consists of three main components:

- **main.go**: The entry point of the program, responsible for creating local server on port 8080 and routing URL paths.
- **createcharacters.go**: A package that reads banner file and creates a 2D slice of ascii art characters.
- **handleln.go**: A package that Handles newline characters in the input string. Prints ASCII art for each character in the input string.
- **printer.go**: A package that prints the ASCII art for each character in the input string.
- **indexhandler.go**: A package that handles the landing page of the web tool
- **formhandler.go**: A package that parses the form values and handles template parsing

### Built With

- Go Programming Language
- HTML

## Getting Started

To get started with ASCII Art Web, follow the instructions below.

### Prerequisites

Before running the program, ensure that you have the following prerequisites:

- Go installed on your machine.
- Basic understanding of Go programming language.

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Philip38-hub/ascii-art-web.git
    ```

2. Navigate to the project directory:

    ```sh
    cd ascii-art-web
    ```

## Usage

To use ASCII Art Web, follow these steps:

1. Prepare a text file containing the text you want to convert into ASCII art.

2. Run the program with the following command:

    ```sh
    go run . 
    ```
3. Open your browser and on a new tab, go to ```localhost:8080 ```
4. Input your text, choose your preferred style of display(banner) and then submit by pressing the generate ascii art button

## Optional ASCII Art Styles

ASCII Art Output supports multiple ASCII art styles, including:

- **Standard**: A basic ASCII art style.
- **Shadow**: ASCII art with shadow effects.
- **Thinkertoy**: ASCII art with a playful design.

## Expected Output

Instance 1. shadow 
```
                                       
_|_|_|_| _|                     _|     
_|          _|  _|_|   _|_|_| _|_|_|_| 
_|_|_|   _| _|_|     _|_|       _|     
_|       _| _|           _|_|   _|     
_|       _| _|       _|_|_|       _|_| 
                                       
                                       
                                      
_|_|_|_|_|                     _|     
    _|       _|_|     _|_|_| _|_|_|_| 
    _|     _|_|_|_| _|_|       _|     
    _|     _|           _|_|   _|     
    _|       _|_|_| _|_|_|       _|_| 
                                      
```                                   

Instance 2. thinktertoy
```

                                    o o         | |               
  0    --  o-o            o         | |  | |   -O-O-      O   o   
 /|   o  o    |            \            -O-O- o | |   o  /   / \  
o |     /   oo              O            | |   -O-O-    /   o O-o 
  |    /      |       o-o  /            -O-O-   | | o  /  o  \    
o-o-o o--o o-o            o              | |   -O-O-  O       o-  
                                                | |               
                                                                

```

Instance 3. standard

```
                                   __             _  _      _     _   __ 
 _   ____    _____                 \ \          _| || |_   | |   (_) / / 
/ | |___ \  |___ /         ______   \ \        |_  __  _| / __)     / /  
| |   __) |   |_ \        |______|   > >        _| || |_  \__ \    / /   
| |  / __/   ___) |                 / /        |_  __  _| (   /   / / _  
|_| |_____| |____/                 /_/           |_||_|    |_|   /_/ (_) 
                                                                         
                                                                        

```
## Roadmap

The following features are planned for future releases:
- Add support for color.
- Add feature to specify text alignment
- Integration with third-party ASCII art libraries.

## Contributing

Contributions to ASCII Art Output are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your fork.
5. Open a pull request to merge your changes into the main branch.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Authors

- **[X - @oumaphilip01](https://x.com/oumaphilip01)**
- **[Github - Philip38](https://github.com/Philip38-hub)**

<p align="right">(<a href="#ascii-art-web">back to top</a>)</p>


## Acknowledgments

Special thanks to the creators of the ASCII art character sets used in this project.

