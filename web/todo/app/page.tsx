"use client"

import TopHeader from "@/app/header/header";
import ChoresPage from "@/app/chores/page";
import {Chore} from "@/app/chores/chores";

export default function Home() {



    return (
        <main className="flex min-h-full flex-col items-center justify-between p-24">
        <TopHeader />

        <div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
        </div>
    </main>
  );
}
