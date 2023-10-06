import { Button } from "@/components/ui/button";
import React, { useState } from "react";
import { GetRandomImageUrl } from "../wailsjs/go/main/App.js";

function App() {
  const [randomImageUrl, setRandomImageUrl] = useState<string>("");
  const fetchRandomImage = async () => {
    const image = await GetRandomImageUrl();
    setRandomImageUrl(image);
  };

  return (
    <div className="min-h-screen bg-white grid place-items-center mx-auto py-8">
      <div className="text-blue-900 text-2xl font-bold flex flex-col items-center space-y-4">
        <h1>Vite + React + TS + Tailwind + shadcn/ui</h1>
        <Button onClick={() => fetchRandomImage()}>fetch RandomImage</Button>
        {randomImageUrl !== "" && (
          <img src={randomImageUrl} className="w-[320px]" alt="" />
        )}
      </div>
    </div>
  );
}

export default App;
