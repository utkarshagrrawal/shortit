import { useEffect, useState } from "react";
import LinkIcon from "../components/icons/link";
import axios from "axios";

export default function Home() {
  const [url, setUrl] = useState("");
  const [response, setResponse] = useState({
    message: "",
    type: "",
  });
  const [loading, setLoading] = useState(false);
  const [urls, setUrls] = useState([]);

  useEffect(() => {
    try {
      setUrls(JSON.parse(localStorage.getItem("urls") || "[]"));
    } catch (error) {
      localStorage.removeItem("urls");
      setUrls([]);
    }
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();

    setLoading(true);

    setResponse({
      message: "",
      type: "",
    });

    axios
      .post(
        import.meta.env.VITE_API_URL + "/api/v1/shorten",
        { url },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      )
      .then((res) => {
        setResponse({
          message: "Link shortened successfully: " + res.data,
          type: "success",
        });
        let oldUrls;
        try {
          oldUrls = JSON.parse(localStorage.getItem("urls") || "[]");
        } catch (error) {
          oldUrls = [{ originalUrl: url, shortUrl: res.data }];
        }
        let isUrlPresent = oldUrls.find((u) => u.originalUrl === url);
        if (!isUrlPresent) {
          setUrls([...oldUrls, { originalUrl: url, shortUrl: res.data }]);
          localStorage.setItem(
            "urls",
            JSON.stringify([
              ...oldUrls,
              { originalUrl: url, shortUrl: res.data },
            ])
          );
        } else {
          setUrls(oldUrls);
        }
      })
      .catch((err) => {
        setResponse({
          message:
            err.response?.data || "An error occurred while shortening the link",
          type: "error",
        });
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleDelete = () => {};

  return (
    <div className="min-h-screen w-full">
      <header className="flex items-center justify-between p-4 bg-white">
        <div className="flex items-center gap-2">
          <LinkIcon className="size-6" />
          <h3 className="text-xl font-bold">Shortit</h3>
        </div>
        <nav className="flex items-center gap-4">
          <a
            href="/analytics"
            className="text-gray-600 hover:text-gray-800 text-sm font-semibold hover:underline"
          >
            Analytics
          </a>
          <a
            href="/privacy"
            className="text-gray-600 hover:text-gray-800 text-sm font-semibold hover:underline"
          >
            Privacy Policy
          </a>
        </nav>
      </header>
      <div className="flex items-center justify-center w-full">
        <div className="bg-white max-w-2xl rounded-lg p-8">
          <h1 className="text-3xl font-extrabold text-center text-gray-800 mb-4">
            Shorten Your Links with{" "}
            <span className="text-blue-500">Shortit</span>
          </h1>
          <p className="text-gray-600 text-center mb-6">
            Shortit is a simple{" "}
            <span className="text-blue-500 font-semibold">URL shortener</span>{" "}
            that helps you turn long links into shorter, shareable ones.
          </p>

          {response.message && (
            <div
              className={`p-2 text-center mb-4 rounded-lg font-semibold text-sm border ${
                response.type === "error"
                  ? "text-red-500 bg-red-100 border-red-500"
                  : "text-green-500 bg-green-100 border-green-500"
              }`}
            >
              {response.message}
            </div>
          )}

          <form onSubmit={handleSubmit} className="flex items-center gap-2">
            <input
              type="url"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder="Paste your link here"
              className="flex-1 p-3 border border-gray-300 rounded-l-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
            <button
              className={`bg-blue-500 hover:bg-blue-600 transition text-white px-5 py-3 rounded-r-lg font-semibold ${
                loading ? "opacity-50 cursor-not-allowed" : ""
              }`}
              disabled={loading}
            >
              Shorten
            </button>
          </form>

          <div className="mt-4">
            {urls.map((u, i) => (
              <div
                key={i}
                className="flex items-center justify-between p-3 border border-gray-200 rounded-lg mt-4"
              >
                <div>
                  <p className="text-gray-600 text-sm">{u.originalUrl}</p>
                  <a
                    href={u.shortUrl}
                    target="_blank"
                    rel="noreferrer"
                    className="text-blue-500 text-sm font-semibold hover:underline"
                  >
                    {u.shortUrl}
                  </a>
                </div>
                <button
                  className="text-red-500 text-sm font-semibold"
                  onClick={handleDelete}
                >
                  Delete
                </button>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
