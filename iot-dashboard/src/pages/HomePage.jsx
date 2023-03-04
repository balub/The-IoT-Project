import React from "react";
import { Link } from "react-router-dom";
import { useState } from "react";
import NavBar from "../components/NavBar";
import Popup from "reactjs-popup";
import AddProject from "../components/AddProject";

export default function HomePage() {
  const [projects, setProjects] = useState([
    { name: "Project 1", id: "y]d?va+9b " },
    { name: "Project 2", id: "5c;H7x27 " },
    { name: "Project 1", id: "y]d?va+9b " },
    { name: "Project 2", id: "5c;H7x27 " },
    { name: "Project 1", id: "y]d?va+9b " },
    { name: "Project 2", id: "5c;H7x27 " },
    { name: "Project 1", id: "y]d?va+9b " },
    { name: "Project 2", id: "5c;H7x27 " },
    { name: "Project 1", id: "y]d?va+9b " },
    { name: "Project 2", id: "5c;H7x27 " },
  ]);

  return (
    <div>
      <NavBar>
        <header className="bg-white shadow-sm">
          <div className="flex justify-between max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8">
            <h1 className="text-lg leading-6 font-semibold text-gray-900">
              My Projects
            </h1>
            <Popup
              trigger={
                <button className="bg-slate-300 p-2 rounded font-bold">
                  + Add Project
                </button>
              }
              position=""
            >
              <AddProject />
            </Popup>
          </div>
        </header>
      </NavBar>
      <div className="mt-6 justify-center flex flex-wrap">
        {projects.map((project) => (
          <Link to="/projects/project">
            <div
              className="bg-indigo-100 my-2 mx-2 w-48 rounded-2xl h-48 p-10"
              onClick={() => handleOnClick}
            >
              <div className="font-bold text-xl">{project.name}</div>
              <div className="text-slate-500">{project.id}</div>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
}
