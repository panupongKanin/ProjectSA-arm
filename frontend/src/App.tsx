import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Users from "./components/test";
import SimplePaper from "./components/MapBedPage";
//import Navbar from "./components/Navbar";

export default function App() {
return (
  <Router>
    <SimplePaper/>
  </Router>
  );
}