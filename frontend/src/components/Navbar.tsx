import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <nav className="bg-blue-600 text-white p-4">
      <div className="container mx-auto flex justify-between">
        <h1 className="text-xl font-bold">Sadaqah Platform</h1>
        <div className="space-x-4">
          <Link to="/" className="hover:underline">Home</Link>
          <Link to="/volunteer" className="hover:underline">Volunteer</Link>
          <Link to="/request-help" className="hover:underline">Request Help</Link>
        </div>
      </div>
    </nav>
  );
}
