"use client"

import {Chore} from "@/app/chores/chores";
import {createClient} from "@hey-api/client-fetch";
import React, {useEffect, useState} from "react";
import {getChores} from "@/app/todoclient";

createClient({
    baseUrl:'http://localhost:8080',
})

function daysBetweenDates(d1: number, d2: number): number {
    return Math.round((d2-d1) / (1000 * 3600 * 34))
}

function timeDeltaDisplay(daysBetween: number): string {
    if (daysBetween <= -1) {
        return `${Math.abs(daysBetween)} days ago`
    }
    if (daysBetween > 0) {
        return `in ${Math.abs(daysBetween)} days`
    }
    return "today"
}

export default function ChoresPage() {
    const [chores, setChores] = useState<Array<Chore>>([
        {
            id: 1,
            name: "Scrub Toilets",
            // next_deadline: new Date().setDate(Date.now()) + 5,
            // last_completed: new Date().setDate(Date.now()) - 2,
            next_deadline: Date.now() + (1000 * 60 * 60 * 24 * 5),
            last_completed: Date.now() - (1000 * 60 * 60 * 24 * 2),
        }
    ])

    useEffect(() => {
        getChores()
            .then((resp) => {
                setChores(resp.chores!)
            })
            .catch(e => {
                console.error(e)
            })
    }, [])


    return (
        <div className="chores-page">
            {chores.map(chore => (
                <div className="chore max-w-sm rounded overflow-hidden shadow-lg" key={chore.id}>
                    <div className={"px-6 py-4"}>
                        <div className={"fond-bold text-xl mb-2"}>{chore.name}</div>
                    </div>
                    <div className={"px-6 pt-4 pb-2"}>
                        <span className={"inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"}>Last Completed: {timeDeltaDisplay(daysBetweenDates(Date.now(), chore.last_completed))}</span>
                        <span className={"inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"}>Next Deadline: {timeDeltaDisplay(daysBetweenDates(Date.now(), chore.next_deadline))}</span>
                    </div>

                </div>
            ))}
        </div>
    )
}
