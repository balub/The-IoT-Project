import React, { useState, useEffect } from "react";

function DataModelInput({ data, addModelData }) {
  const [title, settitle] = useState("");
  const [type, settype] = useState("");
  const [Required, setRequired] = useState("");
  const [Component, setComponent] = useState("");

  useEffect(() => {
    addModelData({
      ...data,
      title,
      type,
      required: Required,
      component: Component,
    });
  }, [title, type, Required, Component]);

  return (
    <div className="flex mt-16">
      <div className="mx-2">
        <label
          htmlFor="email"
          className="block text-sm font-medium text-gray-500"
        >
          Title
        </label>
        <input
          type="text"
          name="title"
          id="title"
          className="p-1 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
          placeholder="ID"
          value={title}
          onChange={(e) => settitle(e.target.value)}
        />
      </div>
      <div className="mx-2">
        <label
          htmlFor="email"
          className="block text-sm font-medium text-gray-500"
        >
          Type
        </label>
        <input
          type="type"
          name="type"
          id="type"
          className="p-1 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
          placeholder="Array"
          value={type}
          onChange={(e) => settype(e.target.value)}
        />
      </div>
      <div className="mx-2">
        <label
          htmlFor="email"
          className="block text-sm font-medium text-gray-500"
        >
          Required
        </label>
        <input
          type="required"
          name="required"
          id="required"
          className="p-1 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
          placeholder="true"
          value={Required}
          onChange={(e) => setRequired(e.target.value)}
        />
      </div>
      <div className="mx-2">
        <label
          htmlFor="email"
          className="block text-sm font-medium text-gray-500"
        >
          Component
        </label>
        <input
          type="component"
          name="component"
          id="component"
          className="p-1 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
          placeholder=""
          value={Component}
          onChange={(e) => setComponent(e.target.value)}
        />
      </div>
    </div>
  );
}

export default DataModelInput;
