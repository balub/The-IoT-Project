import React from "react";
import useTabsStore from "../store/useTabsStore";

function Tab({ tabCount, title, onClick }) {
  const selectedTab = useTabsStore((state) => state.selectedTab);
  console.log(selectedTab, tabCount);

  function classNames(...classes) {
    return classes.filter(Boolean).join(" ");
  }

  return (
    <div
      className={` group relative min-w-0 flex-1 overflow-hidden bg-white py-4 px-4 text-sm font-medium text-center hover:bg-gray-50 focus:z-10`}
      style={{ borderBottom: selectedTab === tabCount ? "2px solid blue" : "" }}
      onClick={onClick}
    >
      {title}
    </div>
  );
}

export default Tab;
