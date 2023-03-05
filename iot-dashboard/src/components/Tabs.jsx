import React from "react";
import useTabsStore from "../store/useTabsStore";
import Tab from "./Tab";

function Tabs() {
  const changeTab = useTabsStore((state) => state.changeTab);

  function onHandleClick(tabNum) {
    changeTab(tabNum);
  }

  return (
    <nav className="relative z-0 rounded-lg shadow flex  divide-x divide-gray-200 ">
      <Tab
        tabCount={1}
        title={"Device Management"}
        onClick={() => onHandleClick(1)}
      />
      <Tab tabCount={2} title={"Data Model"} onClick={() => onHandleClick(2)} />
      <Tab tabCount={3} title={"Api"} onClick={() => onHandleClick(3)} />
    </nav>
  );
}

export default Tabs;
