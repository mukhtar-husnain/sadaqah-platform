import { Link } from "react-router-dom";

export default function Volunteer() {
    return (
      <div className="text-center p-6">
        <h1 className="text-3xl font-bold">Become a Volunteer</h1>
        <p className="text-lg mt-4">Sign up to help others in need.</p>
        <Link to="/register-volunteer">
            <button className="mt-4 bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700">
            Register as a Volunteer
            </button>
        </Link>
      </div>
    );
  }
  