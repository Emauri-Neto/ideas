import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { cn } from "@/lib/utils";
import { ThemeProvider } from "@/components/theme-provider";

export const metadata: Metadata = {
  title: "Ideas",
  description: "Plataforma idealizadora ideias.",
};

const inter = Inter({ subsets: ["latin"]});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={cn("antialiased", inter.className)}>
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
