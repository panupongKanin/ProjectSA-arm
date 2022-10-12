import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MyComponent from "./components/home";
//import Navbar from "./components/Navbar";

export default function App() {
return (
  <Router>
   <MyComponent/>
  </Router>
  );
}