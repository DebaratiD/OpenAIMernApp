import express from "express";
import * as dotenv from 'dotenv'
import cors from 'cors'

import connectDB from "./mongodb/connect.js";
import postRoutes from './Routes/postRoutes.js';
import dalleRoutes from './Routes/dalleRoutes.js';

dotenv.config() //  pull env variables from the dotenv file

const app = express()
app.use(cors())
app.use(express.json({limit:'50mb'}))

//These are API endpoints that can be hooked on to from the frontend side.
app.use('/api/v1/post', postRoutes);
app.use('/api/v1/dalle', dalleRoutes);

app.get('/', async(req,res)=> {
    res.send('Hello from Dall-E')
})

const startServer = async () => {
    //before listening on the port, we want to connect to mongodb

    try {
        connectDB(process.env.MONGODB_URL)
        app.listen(8080, () => console.log('Server started on port http://localhost:8080'))
    } catch (error) {
        console.log(error)
    }
}

startServer();