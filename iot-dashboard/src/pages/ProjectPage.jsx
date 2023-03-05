import React from "react";
import { useParams } from "react-router-dom";
import DeviceManagement from "../components/DeviceManagement";
import DataModel from "../components/DataModel";
import Api from "../components/Api";
import NavBar from "../components/NavBar";
import Tabs from "../components/Tabs";

import useTabsStore from "../store/useTabsStore";

function ProjectPage() {
  const selectedTab = useTabsStore((state) => state.selectedTab);
  const changeTab = useTabsStore((state) => state.changeTab);

  function tabsRender() {
    switch (selectedTab) {
      case 1:
        return <DeviceManagement />;
      case 2:
        return <DataModel />;
      case 3:
        return <Api />;
    }
  }

  const { projectId } = useParams();
  return (
    <div>
      <NavBar>
        <header className="bg-white shadow-sm">
          <div className="flex justify-between max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8">
            <h1 className="text-lg leading-6 font-semibold text-gray-900">
              Project Name
            </h1>
          </div>
        </header>
      </NavBar>
      <div className="border-b border-gray-200">
        <Tabs />
        <div>{tabsRender()}</div>
      </div>
    </div>
  );
}

export default ProjectPage;
