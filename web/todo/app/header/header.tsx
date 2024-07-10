import Image from "next/image";

export default function TopHeader() {
    return <header className="mb-32 grid text-center lg:mb-0 lg:w-full lg:max-w-5xl lg:grid-cols-4 lg:text-left">
        <a
            href="/"
            className="group rounded-lg border border-transparent px-5 py-4 transition-colors hover:border-gray-300 hover:bg-gray-100 hover:dark:border-neutral-700 hover:dark:bg-neutral-800/30"
            rel="noopener noreferrer"
        >
            <Image
                className="relative dark:drop-shadow-[0_0_0.3rem_#ffffff70]"
                src="/Abandonedfactory_lowpoly.png"
                alt="AbandonedFactory logo"
                width={90}
                height={18}
                priority
            />

        </a>

        <a
            href="/"
            className="group rounded-lg border border-transparent px-5 py-4 transition-colors hover:border-gray-300 hover:bg-gray-100 hover:dark:border-neutral-700 hover:dark:bg-neutral-800/30"
            rel="noopener noreferrer"
        >
            <h2 className="mb-3 text-2xl font-semibold">AF TODOs</h2>
        </a>

        <a
            href="/lists"
            className="group rounded-lg border border-transparent px-5 py-4 transition-colors hover:border-gray-300 hover:bg-gray-100 hover:dark:border-neutral-700 hover:dark:bg-neutral-800/30"
            rel="noopener noreferrer"
        >
            <h2 className="mb-3 text-2xl font-semibold">
                Lists{" "}
                <span className="inline-block transition-transform group-hover:translate-x-1 motion-reduce:transform-none">
              -&gt;
            </span>
            </h2>
            <p className="m-0 max-w-[30ch] text-sm opacity-50">
                Manage your TODOs
            </p>
        </a>

        <a
            href="/chores"
            className="group rounded-lg border border-transparent px-5 py-4 transition-colors hover:border-gray-300 hover:bg-gray-100 hover:dark:border-neutral-700 hover:dark:bg-neutral-800/30"
            rel="noopener noreferrer"
        >
            <h2 className="mb-3 text-2xl font-semibold">
                Chores{" "}
                <span className="inline-block transition-transform group-hover:translate-x-1 motion-reduce:transform-none">
              -&gt;
            </span>
            </h2>
            <p className="m-0 max-w-[30ch] text-sm opacity-50">
                Chores
            </p>
        </a>

        <a
            href="/login"
            className="group rounded-lg border border-transparent px-5 py-4 transition-colors hover:border-gray-300 hover:bg-gray-100 hover:dark:border-neutral-700 hover:dark:bg-neutral-800/30"
            rel="noopener noreferrer"
        >
            <h2 className="mb-3 text-2xl font-semibold">
                Login{" "}
                <span className="inline-block transition-transform group-hover:translate-x-1 motion-reduce:transform-none">
              -&gt;
            </span>
            </h2>
        </a>

    </header>
}