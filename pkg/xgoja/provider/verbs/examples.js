__package__({ name: "examples" })

__verb__("typedValues", {
  name: "typed-values"
})
function typedValues() {
  const dbus = require("dbus")
  const count = dbus.u32(42)
  const path = dbus.path("/com/example/App1")
  const variant = dbus.variant("s", "hello")
  const actions = dbus.array("as", ["default", "Open"])
  const hints = dbus.dict("a{sv}", {
    urgency: dbus.variant("u", dbus.u32(1))
  })
  const pair = dbus.struct("(su)", ["count", dbus.u32(7)])

  return {
    countSignature: count.signature,
    countValue: count.value,
    pathSignature: path.signature,
    pathValue: path.value,
    variantSignature: variant.signature,
    variantInnerSignature: variant.value.signature,
    actionsSignature: actions.signature,
    actionsCount: actions.value.length,
    hintsSignature: hints.signature,
    hintUrgencySignature: hints.value.urgency.signature,
    pairSignature: pair.signature,
    pairFirst: pair.value[0]
  }
}

__verb__("deniedSystemBus", {
  name: "denied-system-bus"
})
async function deniedSystemBus() {
  const dbus = require("dbus")
  try {
    await dbus.system().connect()
    return { denied: false, code: "", message: "unexpectedly connected" }
  } catch (err) {
    return {
      denied: true,
      code: err.code || "",
      name: err.name || "",
      message: String(err.message || err)
    }
  }
}

__verb__("getIdScript", {
  name: "get-id-script",
  output: "text"
})
function getIdScript() {
  return `const dbus = require("dbus");

const bus = await dbus.session().timeout(2000).connect();
try {
  const id = await bus
    .destination("org.freedesktop.DBus")
    .object("/org/freedesktop/DBus")
    .interface("org.freedesktop.DBus")
    .method("GetId")
    .out("s")
    .call();
  console.log("bus id:", id);
} finally {
  await bus.close();
}
`;
}
