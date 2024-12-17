export default function PrivacyPolicy() {
  return (
    <div className="p-4">
      <a
        href="/"
        className="text-black hover:bg-black hover:text-white duration-200 border border-gray-500 rounded-lg px-4 py-2"
      >
        Back to Home
      </a>
      <section className="px-4 sm:px-6 lg:px-8">
        <div className="max-w-4xl mx-auto">
          <h1 className="text-4xl font-bold text-center mb-8 text-gray-900">
            Privacy Policy
          </h1>

          <p className="text-gray-700 text-lg mb-6">
            Your privacy is important to us. This Privacy Policy explains how we
            collect, use, and protect your information when you use our URL
            shortening service.
          </p>

          {/* Section 1: Data Collection */}
          <div className="mb-8">
            <h2 className="text-2xl font-semibold text-gray-900 mb-4">
              1. What Information We Collect
            </h2>
            <ul className="list-disc pl-6 text-gray-700 space-y-2">
              <li>
                <strong>IP Address:</strong> We collect the IP address of users
                when they create a shortened URL.
              </li>
              <li>
                <strong>Device Information:</strong> We collect device details
                like browser type, operating system, and user-agent information.
              </li>
              <li>
                <strong>Timestamps:</strong> We store the date and time when the
                URL was created.
              </li>
            </ul>
          </div>

          {/* Section 2: Purpose of Collection */}
          <div className="mb-8">
            <h2 className="text-2xl font-semibold text-gray-900 mb-4">
              2. Why We Collect This Information
            </h2>
            <p className="text-gray-700 mb-4">
              We collect this data for the following reasons:
            </p>
            <ul className="list-disc pl-6 text-gray-700 space-y-2">
              <li>
                To prevent misuse of our service, such as spam, phishing, and
                other malicious activity.
              </li>
              <li>
                To comply with legal obligations and assist law enforcement if
                required.
              </li>
              <li>To ensure a safe and secure experience for all users.</li>
            </ul>
          </div>

          {/* Section 3: Data Protection */}
          <div className="mb-8">
            <h2 className="text-2xl font-semibold text-gray-900 mb-4">
              3. How We Protect Your Data
            </h2>
            <p className="text-gray-700 mb-4">
              We take your privacy seriously and implement the following
              safeguards:
            </p>
            <ul className="list-disc pl-6 text-gray-700 space-y-2">
              <li>
                Access controls to restrict who can view sensitive information.
              </li>
              <li>
                Regular security audits to protect against unauthorized access
                or data breaches.
              </li>
            </ul>
          </div>

          {/* Section 4: User Rights */}
          <div className="mb-8">
            <h2 className="text-2xl font-semibold text-gray-900 mb-4">
              4. Your Rights
            </h2>
            <p className="text-gray-700 mb-4">
              As a user, you have the right to:
            </p>
            <ul className="list-disc pl-6 text-gray-700 space-y-2">
              <li>
                Request that we delete your data (including your IP) from our
                system.
              </li>
              <li>
                Request information on what data we have collected about you.
              </li>
              <li>Opt-out of tracking, where legally required.</li>
            </ul>
          </div>

          {/* Section 5: Contact Us */}
          <div className="mb-8">
            <h2 className="text-2xl font-semibold text-gray-900 mb-4">
              5. Contact Us
            </h2>
            <p className="text-gray-700 mb-4">
              If you have any questions about this Privacy Policy or wish to
              request access to your data, please contact us at:
            </p>
            <p className="text-gray-900 font-medium">
              <a
                href="mailto:utkarsh09jan@gmail.com"
                className="text-blue-500 hover:underline"
              >
                utkarsh09jan@gmail.com
              </a>
            </p>
          </div>

          {/* Call-to-action or footer */}
          <div className="text-center mt-12">
            <p className="text-gray-700">
              This Privacy Policy may be updated periodically. Please review it
              frequently to stay informed.
            </p>
          </div>
        </div>
      </section>
    </div>
  );
}
