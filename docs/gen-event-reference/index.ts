import { events } from 'web/packages/teleport/src/Audit/fixtures';
import { makeEvent } from 'web/packages/teleport/src/services/audit/makeEvent';

events.forEach(e => {
  console.log(makeEvent(e));
});
