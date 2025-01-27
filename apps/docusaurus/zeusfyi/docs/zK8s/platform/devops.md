---
sidebar_position: 5
displayed_sidebar: zK8s
---

# DevOps

## Dashboard

This panel allows you to perform DevOps tasks on your cluster, namespace, and workloads. It's also ideal for
offloading technical support tasks to your customers, or for your own internal support team.

![Panel](https://github.com/zeus-fyi/zeus/assets/17446735/a20885e5-07a1-47fd-bf53-b31283a7b579)

## Workload Actions

### zK8s App Page

This will take you to the zK8s App Page for your workload. This is a great way to share your workload with
your customers or internal support team, or for pair programming. It's also extremely useful
for rapid prototype development and debugging, since you can change & redeploy your workload in real-time.

### Deploy Latest

This will deploy the latest version of your workload. It will also restart your workload.

### Rollout Restart

This will restart your workloads in a rolling fashion. If you have a docker image with a latest tag, and
you have your ImagePullPolicy set to Always, it'll pull down the latest image and restart your workload
for it to take effect. This is extremely useful for rapid debugging & development.

## Pods Actions

### Get Logs

This will dump the logs from your pod. You can also specify the container name if you have multiple containers

### Delete Pod

This will delete your pod. It will be recreated automatically if you're using a deployment or statefulset, or replica
set. 