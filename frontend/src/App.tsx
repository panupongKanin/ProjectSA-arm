import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import HomePage from "./components/home";
//import Navbar from "./components/Navbar";

export default function App() {
return (
  <Router>
   <HomePage/>
  </Router>
  );
}