
import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MyComponent from "./components/login";
import SimpleContainer from "./components/MapBedPage";
import SimplePaper from "./components/MapBedPage";

//import Navbar from "./components/Navbar";

export default function App() {
return (
  <Router>
    <SimpleContainer/>
    {/* <MyComponent/> */}
  </Router>
  );
}