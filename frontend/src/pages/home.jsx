import { useEffect, useState } from "react";
import LinkIcon from "../components/icons/link";
import axios from "axios";
import { useParams } from "react-router-dom";
import ClickIcon from "../components/icons/click";

export default function Home() {
  const { shortUrl } = useParams();
  const [url, setUrl] = useState("");
  const [response, setResponse] = useState({
    message: "",
    type: "",
  });
  const [loading, setLoading] = useState(false);
  const [redirecting, setRedirecting] = useState(false);
  const [urls, setUrls] = useState([]);
  const [ip, setIp] = useState("");

  useEffect(() => {
    if (shortUrl) {
      setRedirecting(true);
      axios
        .get(import.meta.env.VITE_API_URL + "/api/v1/" + shortUrl)
        .then((res) => {
          window.location.href = res.data;
        })
        .catch((err) => {
          setRedirecting(false);
          setResponse({
            message:
              err.response?.data ||
              "An error occurred while expanding the link",
            type: "error",
          });
        });
    }
  }, []);

  useEffect(() => {
    axios
      .get("https://api.ipify.org")
      .then((res) => {
        setIp(res.data);
      })
      .catch((err) => {
        console.log("Error fetching IP address");
      });
  }, []);

  useEffect(() => {
    axios
      .get(import.meta.env.VITE_API_URL + "/api/v1/user/urls", {
        withCredentials: true,
      })
      .then((res) => {
        setUrls(res.data || []);
      })
      .catch((err) => {
        console.log("Error fetching user URLs");
      });
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
        { url, ip },
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true,
        }
      )
      .then((res) => {
        setResponse({
          message: "Link shortened successfully: " + res.data,
          type: "success",
        });
        let isUrlPresent =
          urls.length > 0 && urls.find((u) => u.shortUrl === res.data);
        if (!isUrlPresent) {
          setUrls([
            ...urls,
            {
              originalUrl: url,
              shortUrl: res.data,
              clicks: 0,
              createdAt: new Date().toISOString(),
            },
          ]);
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

  const handleDelete = (url) => {
    const shortId = url.shortUrl.split("/");
    axios
      .delete(import.meta.env.VITE_API_URL + "/api/v1/" + shortId[3], {
        withCredentials: true,
      })
      .then(() => {
        setUrls(urls.filter((u) => u.shortUrl !== url.shortUrl));
      })
      .catch((err) => {
        setResponse({
          message:
            err.response?.data || "An error occurred while deleting the link",
          type: "error",
        });
      });
  };

  return redirecting ? (
    <div className="min-h-screen w-full flex items-center justify-center">
      <div className="grid min-h-[140px] w-full place-items-center overflow-x-scroll rounded-lg p-6 lg:overflow-visible">
        <svg
          className="w-16 h-16 animate-spin text-gray-900/50"
          viewBox="0 0 64 64"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
        >
          <path
            d="M32 3C35.8083 3 39.5794 3.75011 43.0978 5.20749C46.6163 6.66488 49.8132 8.80101 52.5061 11.4939C55.199 14.1868 57.3351 17.3837 58.7925 20.9022C60.2499 24.4206 61 28.1917 61 32C61 35.8083 60.2499 39.5794 58.7925 43.0978C57.3351 46.6163 55.199 49.8132 52.5061 52.5061C49.8132 55.199 46.6163 57.3351 43.0978 58.7925C39.5794 60.2499 35.8083 61 32 61C28.1917 61 24.4206 60.2499 20.9022 58.7925C17.3837 57.3351 14.1868 55.199 11.4939 52.5061C8.801 49.8132 6.66487 46.6163 5.20749 43.0978C3.7501 39.5794 3 35.8083 3 32C3 28.1917 3.75011 24.4206 5.2075 20.9022C6.66489 17.3837 8.80101 14.1868 11.4939 11.4939C14.1868 8.80099 17.3838 6.66487 20.9022 5.20749C24.4206 3.7501 28.1917 3 32 3L32 3Z"
            stroke="currentColor"
            strokeWidth="5"
            strokeLinecap="round"
            strokeLinejoin="round"
          ></path>
          <path
            d="M32 3C36.5778 3 41.0906 4.08374 45.1692 6.16256C49.2477 8.24138 52.7762 11.2562 55.466 14.9605C58.1558 18.6647 59.9304 22.9531 60.6448 27.4748C61.3591 31.9965 60.9928 36.6232 59.5759 40.9762"
            stroke="currentColor"
            strokeWidth="5"
            strokeLinecap="round"
            strokeLinejoin="round"
            className="text-gray-900"
          ></path>
        </svg>
      </div>
    </div>
  ) : (
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

          <div className="mt-4 max-h-56 overflow-auto">
            {urls.map((u, i) => (
              <div
                key={i}
                className="flex flex-col sm:flex-row items-center justify-between p-4 bg-white shadow-md border border-gray-200 rounded-lg mt-4 hover:shadow-lg transition-shadow"
              >
                {/* URL Information Section */}
                <div className="flex-1">
                  <p className="text-gray-600 text-sm mb-2 break-words">
                    <span className="font-medium text-gray-800">
                      Original URL:
                    </span>{" "}
                    {u.originalUrl}
                  </p>
                  <a
                    href={u.shortUrl}
                    target="_blank"
                    rel="noreferrer"
                    className="text-blue-500 text-sm font-semibold hover:underline break-words"
                  >
                    {u.shortUrl}
                  </a>
                </div>

                {/* Click Count Section */}
                <div className="flex items-center mt-4 sm:mt-0 sm:ml-6">
                  <div className="flex items-center bg-blue-100 text-blue-600 font-semibold text-sm px-4 py-2 rounded-lg mr-4">
                    <ClickIcon className="size-5 mr-1" />
                    {u.clicks || 0} Clicks
                  </div>

                  {/* Delete Button */}
                  <button
                    className="bg-red-500 text-white text-sm font-semibold px-4 py-2 rounded-lg hover:bg-red-600 transition-colors"
                    onClick={() => handleDelete(u)}
                  >
                    Delete
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
