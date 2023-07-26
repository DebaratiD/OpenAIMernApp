import express from "express";
import * as dotenv from 'dotenv'
import cors from 'cors'

dotenv.config() //  pull env variables from the dotenv file

const app = express()
app.use(cors())
app.use(express.json({limit:'50mb'}))

1:02:33