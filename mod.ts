import { getEnsText } from "https://esm.sh/viem@2.13.3/ens";
import { mainnet } from "https://esm.sh/viem@2.13.3/chains";
import { createPublicClient, http } from "https://esm.sh/viem@2.13.3";

const publicClient = createPublicClient({
  transport: http("https://eth.llamarpc.com"),
  chain: mainnet,
});

const errorMsg = "Invalid module format";

async function handler(req: Request): Promise<Response> {
  const url = new URL(req.url);

  const [name, versionAndPath] = url.pathname.slice(1).split("@");
  if (!name || !versionAndPath) return new Response(errorMsg, { status: 400 });
  const [version, path] = versionAndPath.split("/");

  const swarmCid = await getEnsText(publicClient, {
    name: `${name}.swarmpm.eth`,
    key: version,
  });

  const res = await fetch(
    `https://api.gateway.ethswarm.org/bzz/${swarmCid}/${path}`,
  );

  if (!res.ok) {
    const json = await res.json();
    return new Response(`Failed to fetch from Swarm: ${json.message}`);
  }

  const text = await res.text();

  return new Response(
    text,
    {
      headers: res.headers,
    },
  );
}

Deno.serve(handler);
