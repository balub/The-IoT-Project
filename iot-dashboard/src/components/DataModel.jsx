import React from "react";
import DataModelInput from "./DataModelInput";
import useDataModelStore from "./store/useDataModelStore";

function DataModel() {
  const models = useDataModelStore((state) => state.models);
  const addElementToArray = useDataModelStore(
    (state) => state.addElementToArray
  );
  const updateModelState = useDataModelStore((state) => state.updateModelState);

  const addModelData = (fieldData) => {
    const arrayWithoutItem = models.filter((item) => item.id !== fieldData.id);
    updateModelState([...arrayWithoutItem, { ...fieldData }]);
  };

  return (
    <div className="p-10">
      <div>
        <div>
          {models.map((item) => (
            <DataModelInput
              key={item.id}
              data={item}
              addModelData={addModelData}
            />
          ))}
        </div>
        <button
          className="mt-10 font-bold text-green-500"
          onClick={() =>
            addElementToArray({
              title: "jhfvb",
              type: "dfvdfv",
              required: true,
              component: "dfvdv",
            })
          }
        >
          + Add field
        </button>
      </div>
    </div>
  );
}

export default DataModel;
