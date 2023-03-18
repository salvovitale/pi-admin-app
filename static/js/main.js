async function reboot() {
  const response = await fetch('/api/reboot', { method: 'POST' });
  if (response.ok) {
    alert('Raspberry Pi will reboot shortly.');
  } else {
    alert('Failed to reboot Raspberry Pi');
  }
}
