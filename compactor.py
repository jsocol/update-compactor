def find_updates_from_msg(msg):
    updates = {}
    for path in msg.update_mask.paths:
        curr = msg.person
        steps = path.split('.')
        leaf = steps.pop()
        for step in steps:
            curr = curr.__getattribute__(step)
        updates[path] = curr.__getattribute__(leaf)
    return updates


def update_entity_path(entity, path, value):
    curr = entity
    steps = path.split('.')
    leaf = steps.pop()
    for step in steps:
        curr = curr[step]
    curr[leaf] = value


def update_entity(entity, updates):
    for path, value in updates.items():
        update_entity_path(entity, path, value)
