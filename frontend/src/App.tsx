
import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MyComponent from "./components/login";

import MappingBedCreate from "./components/MappingBed";

//import Navbar from "./components/Navbar";

export default function App() {
return (
  <Router>
    <MappingBedCreate/>
    {/* <MyComponent/> */}
  </Router>
  );
}