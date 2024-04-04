import React from 'react'
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom'
import { painter }  from './assets'
import { Home, CreatePost } from './pages'

const App = () => {
  return (
  <BrowserRouter>
    <header id="headerClass" className="w-full flex justify-between items-center sm:px-8 px-4 py-3">
      <Link to="/">
        <img src={painter} alt='logo' className="w-16 object-contain"/>
      </Link>
      <Link to="/create-post" className="font-inter font-medium bg-[#6469ff] text-white px-4 py-2 rounded-md">Create</Link>
    </header>
    <main id="mainClass" className="sm:p-8 px-4 py-8 w-full  min-h-[calc(100vh-73px)]">
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/create-post" element={<CreatePost />} />
      </Routes>
    </main>
  </BrowserRouter>
  )
}

export default App