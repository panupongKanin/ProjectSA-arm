import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
//import Navbar from "./components/Navbar";
import home from "./components/home";
import UserCreate from "./components/login";
export default function App() {
return (
  <Router>
   <div>
   <Navbar />
   <Routes>
       <Route path="/" element={<Users />} />
       <Route path="/create" element={<UserCreate />} />
   </Routes>
   </div>
  </Router>
  );
}