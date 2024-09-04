import { events } from 'teleport/src/Audit/fixtures';
import { makeEvent } from 'teleport/src/services/audit/makeEvent';

events.forEach(e => {
  console.log(makeEvent(e));
});
