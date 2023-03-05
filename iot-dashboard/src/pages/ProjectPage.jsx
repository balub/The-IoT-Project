import React from "react";
import { useParams } from "react-router-dom";

function ProjectPage() {
  const { projectId } = useParams();

  return (
    <div>
      ProjectPage
      <h1>{projectId}</h1>
    </div>
  );
}

export default ProjectPage;


