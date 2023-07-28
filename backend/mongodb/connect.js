import mongoose, { connect } from 'mongoose'

const connectDB = (url) => {
    mongoose.set('strictQuery', true)

    //Connect to database
    mongoose.connect(url)
    .then(() => console.log('MongoDB connected'))
    .catch((err)=>console.log(err))
}

export default connectDB;