import type { Metadata } from "next";
import { Roboto } from "next/font/google";
import "./globals.css";
import { cn } from "@/lib/utils";
import { ThemeProvider } from "@/components/theme-provider";

export const metadata: Metadata = {
  title: "Ideas",
  description: "Plataforma idealizadora ideias.",
};

const roboto = Roboto({ subsets: ["latin"], weight: "400" });

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={cn("antialiased", roboto.className)}>
      <body className={"min-h-screen antialiased"}>
        <ThemeProvider
          attribute="class"
          defaultTheme="light"
          enableSystem
          disableTransitionOnChange
        >
          <main className="container mx-auto h-full">{children}</main>
        </ThemeProvider>
      </body>
    </html>
  );
}
