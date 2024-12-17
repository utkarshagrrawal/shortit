import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/home";
import PrivacyPolicy from "./pages/privacyPolicy";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/privacy" element={<PrivacyPolicy />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
