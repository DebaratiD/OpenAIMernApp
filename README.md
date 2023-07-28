# OpenAIMernApp - MERN Application that integrates OpenAI 
This is a project that is built using the four key frameworks of MERN stack: MongoDB, Express, Node and React. Additionally, the application also implements OpenAI as an added tool to generate Images based on some text input that can be specific according to user's interest or randomly generated.

This is a MidJourney and Dall-E clone app.

Cloudinary 

## Steps to creating the app

Run the command ``` npm install vite@latest``` to create an application folder structure from vite. Choose react and then Javascript to get started.


### Install Tailwind CSS

Set up tailwindcss using the following commands:
```bash
npm install -D tailwindcss postcss autoprefixer

npx tailwindcss init -p
```

The second command creates the tailwind css and post css config files. 
Install the Tailwind CSS extension for Visual Studio Code to use the intellisense while coding. In the Tailwind config file add the specific styles for the application and replace the assets folder from the github gist link as provided in the tutorial.

Add 2 more folders: components and pages. Components will hold the code for reusable components while pages will hold the page layouts being shown on the app. There is another folder called constants defined that includes an index.js file with a list of prompts. These can be used to choose a random prompt in-case the user isn't sure about how to provide a text prompt to generate an image.


### Routing

In App.jsx, install 'react-router-dom' that will help implement routing from the main page. Use BrowserRouter to create Links and Routes to the Home page and create post page. Link allows to add links/routes to elements on the page, while Routes defines the routes available on the app.

I have installed the arrow function snippets extension that creates different specifications of arrow function. For example, afee creates an empty arrow function.
In index.css add a link to import specific font styles. 


## Backend

Use  nodemon to start the backend project in start script. The following command installs the dependencies:
```bash
npm install cloudinary cors dotenv express mongoose nodemon openai
```
Specifying "type":"module" in package.json below the description helps use the same import export statements as used in react.

After adding the code for MongoDB connection in index.js and restarting the server (1:07:17), the server crashes since there is no Mongo DB Atlas URL added in the process.env.MONGODB_URL variable.

Go to MongoDB atlas site and follow the instructions to create a new database for free. While connecting to the database, select the drivers method to connect to mongodb via the application. The full atlas string will be available  which you can save in a .env file.


PostRoutes is for creating and retrieving the posts. DalleRoutes is for generating data from the OpenAI API.

For OpenAI, since my credits had expired I had to use another email and phone number to login and generate a new API key. Till 1:28:00, the project is now able to generate images based on the prompt provided.

Why use Cloudinary? Since it will host and store the images generated and make the image retrieval process faster.


Link for the youtube tutorial is [here](https://www.youtube.com/watch?v=EyIvuigqDoA).

Like "rfc", used "rafce" to generate a react component.