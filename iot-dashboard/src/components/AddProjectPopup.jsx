import React from "react";
import { useState } from "react";
import { createProject } from "../services/routes";
// import NewProjectCard from "../components/NewProjectCard";

function AddProjectPopup({ cards }) {
  const [name, setName] = useState("h");
  const [dbUrl, setDbUrl] = useState("k");
  const [dbAuthKey, setDbAuthKey] = useState("hh");
  const [bucketName, setBucketName] = useState("jj");
  const [dbProjectName, setDbProjectName] = useState("jjj");

  async function handleOnClick(event) {
    event.preventDefault();
    if (!!name && !!dbUrl && !!dbAuthKey && !!bucketName && !!dbProjectName) {
      try {
        const result = await createProject(
          name,
          dbUrl,
          dbAuthKey,
          bucketName,
          dbProjectName
        );
      } catch (err) {
        console.log(err);
      }
    }
    // cards.push(result)
  }

  return (
    <>
      <div className=" min-h-full flex flex-col justify-center py-12 sm:px-6 lg:px-8">
        <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
          <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
            <form
              onSubmit={(e) => {
                handleOnClick(e);
              }}
              className="space-y-6"
              action="#"
              method="POST"
            >
              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Project Name
                </label>
                <div className="mt-1">
                  <input
                    id="text"
                    name="email"
                    type="text"
                    autoComplete="email"
                    required
                    className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    onChange={(event) => setName(event.target.value)}
                  />
                </div>
              </div>

              <div>
                <label
                  htmlFor="password"
                  className="block text-sm font-medium text-gray-700"
                >
                  Database URL
                </label>
                <div className="mt-1">
                  <input
                    id="password"
                    name="password"
                    type="password"
                    autoComplete="current-password"
                    required
                    className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    onChange={(event) => setDbUrl(event.target.value)}
                  />
                </div>
              </div>

              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                Database Authorization Key
              </label>
              <div className="mt-1">
                <input
                  id="dbAuthKey"
                  name="dbAuthKey"
                  type="dbAuthKey"
                  autoComplete="dbAuthKey"
                  required
                  className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  onChange={(event) => setDbAuthKey(event.target.value)}
                />
              </div>

              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                Bucket Name
              </label>
              <div className="mt-1">
                <input
                  id="bucketName"
                  name="bucketName"
                  type="bucketName"
                  autoComplete="bucketName"
                  required
                  className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  onChange={(event) => setBucketName(event.target.value)}
                />
              </div>

              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                Database Project Name
              </label>
              <div className="mt-1">
                <input
                  id="dbProjectName"
                  name="dbProjectName"
                  type="dbProjectName"
                  autoComplete="dbProjectName"
                  required
                  className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  onChange={(event) => setDbProjectName(event.target.value)}
                />
              </div>

              <div>
                <button
                  type="submit"
                  className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                  Submit
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </>
  );
}

export default AddProjectPopup;
