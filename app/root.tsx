import type { MetaFunction, LinksFunction } from "@remix-run/node";
import { Provider } from "react-redux";
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "@remix-run/react";
import { Header } from "@/components/UI/Header";
import { store } from '@/store'
import styles from "@/styles/global.css";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Coffe delivery",
  viewport: "width=device-width,initial-scale=1",
});

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: styles },
  { rel: "preconnect", href: "https://fonts.googleapis.com" },
  { rel: "preconnect", href: "https://fonts.gstatic.com" },
  {
    rel: "stylesheet",
    href: "https://fonts.googleapis.com/css2?family=Baloo+2:wght@700;800&family=Roboto:wght@400;700&display=swap",
  },
];

export function ErrorBoundary({ error }: { error: Error }) {
  console.error(error);
  return (
    <html>
      <head>
        <title>Oh no!</title>
        <Meta />
        <Links />
      </head>
      <body>
        <h1>Deu erro</h1>
        <p>{error.message}</p>
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return (
    <html lang="pt-BR">
      <head>
        <Meta />
        <Links />
      </head>
      <body>
        <Provider store={store}>
          <Header />
          <Outlet />
          <ScrollRestoration />
          <Scripts />
          <LiveReload />
        </Provider>
      </body>
    </html>
  );
}
